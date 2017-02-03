package yolmolabs.getstrong.fragments;

import android.content.Context;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.DatePicker;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.Spinner;
import android.widget.TextView;

import com.afollestad.materialdialogs.DialogAction;
import com.afollestad.materialdialogs.MaterialDialog;

import org.json.JSONArray;
import org.json.JSONObject;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.concurrent.Callable;

import bolts.Continuation;
import bolts.Task;
import go.fandroid.Fandroid;
import yolmolabs.getstrong.JSONAdapter;
import yolmolabs.getstrong.AppUtil;
import yolmolabs.getstrong.R;

/**
 * Created by hemantasapkota on 26/06/16.
 */
public class TackFragment extends Fragment {

    private ListView list;

    @Override
    public void onActivityCreated(@Nullable Bundle savedInstanceState) {
        super.onActivityCreated(savedInstanceState);
    }

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        getActivity().setTitle("Track meals, workouts & sleep");
    }

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_track, container, false);

        list = (ListView) v.findViewById(R.id.listView);

        list.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                JSONObject jo = (JSONObject) parent.getItemAtPosition(position);

                try {
                    showDialog(getContext(), jo.getString("timestamp"), jo.getString("description"), jo.getString("value"));
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        });

        loadData();

        return v;
    }

    public void loadData() {
         Task.callInBackground(new Callable<JSONAdapter>() {
            @Override
            public JSONAdapter call() throws Exception {
                ArrayList<JSONObject> jo = new ArrayList<JSONObject>();
                JSONArray jarr = null;
                try {
                    byte[] data = Fandroid.getRecords();
                    String json = new String(data, "UTF-8");
                    jarr = new JSONArray(json);
                    for(int i = 0; i < jarr.length(); i++) {
                        jo.add(jarr.getJSONObject(i));
                    }
                } catch (Exception e) {
                    e.printStackTrace();
                }

                JSONAdapter ja = new JSONAdapter(getContext(), jo) {

                    @Override
                    public int getLayoutID() {
                        return R.layout.fragment_track_item;
                    }

                    @Override
                    public void bind(JSONObject jo, int position, View convertView, ViewGroup parent) {
                        final TextView date = (TextView) convertView.findViewById(R.id.txtDate);
                        final TextView desc = (TextView) convertView.findViewById(R.id.txtDescription);
                        final TextView value = (TextView) convertView.findViewById(R.id.txtValue);
                        final TextView total = (TextView) convertView.findViewById(R.id.txtTotal);

                        View groupView = convertView.findViewById(R.id.groupView);

                        try {
                            final String g = jo.getString("group");
                            if (g.trim().isEmpty()) {
                                groupView.setVisibility(View.GONE);
                            } else {
                                groupView.setVisibility(View.VISIBLE);
                                date.setText(jo.getString("group"));
                            }

                            desc.setText(jo.getString("description"));
                            value.setText(jo.getString("value"));

                            total.setText(Fandroid.totalCaloriesByGroup(g));

                        } catch (Exception e) {
                            e.printStackTrace();
                        }
                    }
                };

                return ja;
            }
        }).continueWith(new Continuation<JSONAdapter, Object>() {
            @Override
            public Object then(Task<JSONAdapter> task) throws Exception {
                list.setAdapter(task.getResult());
                return null;
            }
        }, Task.UI_THREAD_EXECUTOR);
    }

    public void showDialog(final Context context, final String id, final String desc, final String value) {
        String negativeText = "Cancel";
        if (!id.isEmpty()) {
            negativeText = "Delete";
        }

        final String neg = negativeText;

        MaterialDialog md = new MaterialDialog.Builder(context)
                .title("Track here")
                .customView(R.layout.dialog_addtrack, true)
                .positiveText("OK")
                .negativeText(negativeText)
                .onNegative(new MaterialDialog.SingleButtonCallback() {
                    @Override
                    public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                        if (neg.equals("Delete")) {

                            Task.callInBackground(new Callable<Object>() {
                                @Override
                                public Object call() throws Exception {
                                    Fandroid.deleteRecord(id);
                                    return null;
                                }
                            }).continueWith(new Continuation<Object, Object>() {
                                @Override
                                public Object then(Task<Object> task) throws Exception {
                                    loadData();
                                    return null;
                                }
                            }, Task.UI_THREAD_EXECUTOR);

                        }
                    }
                })
                .onPositive(new MaterialDialog.SingleButtonCallback() {
                    @Override
                    public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                        final MaterialDialog d = dialog;
                        // Call in background
                        Task.callInBackground(new Callable<Object>() {
                            @Override
                            public Object call() throws Exception {
                                EditText descTxt = (EditText) d.getCustomView().findViewById(R.id.txtDescription);
                                EditText valueTxt = (EditText) d.getCustomView().findViewById(R.id.txtValue);
                                Spinner valueUnitList = (Spinner) d.getCustomView().findViewById(R.id.listUnits);
                                final DatePicker dp = (DatePicker) d.getCustomView().findViewById(R.id.datePicker);

                                String description = descTxt.getText().toString();
                                String calories = valueTxt.getText().toString();
                                String unit = valueUnitList.getSelectedItem().toString();

                                Date date = AppUtil.getDateFromDatePicker(dp);
                                SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssXXX");
                                String timestamp = sdf.format(date);

                                if (id.isEmpty()) {
                                    Fandroid.addNewRecord(timestamp, description, calories, unit);
                                } else {
                                    Fandroid.updateRecord(id, timestamp, description, calories);
                                }

                                return null;
                            }
                        }).continueWith(new Continuation<Object, Object>() {
                            @Override
                            public Object then(Task<Object> task) throws Exception {
                                Exception e = task.getError();
                                if ( e != null) {
                                    AppUtil.showError(context, e.getMessage()).onAny(new MaterialDialog.SingleButtonCallback() {
                                        @Override
                                        public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                                            showDialog(context, id, desc, value);
                                        }
                                    }).show();
                                } else {
                                    loadData();
                                }
                                return null;
                            }
                        }, Task.UI_THREAD_EXECUTOR);

                    }
                }).show();

        View d = md.getCustomView();

        EditText descTxt = (EditText) d.findViewById(R.id.txtDescription);
        EditText valueTxt = (EditText) d.findViewById(R.id.txtValue);

        final Spinner listUnits = (Spinner) d.findViewById(R.id.listUnits);

        final ArrayAdapter<String> units = new ArrayAdapter<String>(context, android.R.layout.simple_list_item_1);

        Task.callInBackground(new Callable<Object>() {
            @Override
            public Object call() throws Exception {

                byte[] data = Fandroid.getUnits();
                String unitsJson = new String(data, "UTF-8");

                JSONObject jo = new JSONObject(unitsJson);
                JSONArray jarr = jo.getJSONArray("units");

                for(int i = 0; i < jarr.length(); i++) {
                    units.add(jarr.get(i).toString());
                }

                return null;
            }
        }).continueWith(new Continuation<Object, Object>() {
            @Override
            public Object then(Task<Object> task) throws Exception {
                listUnits.setAdapter(units);
                return null;
            }
        }, Task.UI_THREAD_EXECUTOR);

        descTxt.setText(desc);
        valueTxt.setText(value);
    }

}
