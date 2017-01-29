package yolmolabs.getstrong.fragments;

import android.content.Context;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.support.v7.widget.PopupMenu;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.Spinner;
import android.widget.SpinnerAdapter;
import android.widget.TextView;

import com.afollestad.materialdialogs.DialogAction;
import com.afollestad.materialdialogs.MaterialDialog;

import org.json.JSONArray;
import org.json.JSONObject;

import java.util.ArrayList;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;

import bolts.Continuation;
import bolts.Task;
import go.fandroid.Fandroid;
import io.nlopez.smartadapters.SmartAdapter;
import yolmolabs.getstrong.JSONAdapter;
import yolmolabs.getstrong.MainActivity;
import yolmolabs.getstrong.MessageUtil;
import yolmolabs.getstrong.R;

/**
 * Created by hemantasapkota on 26/06/16.
 */
public class MealFragment extends Fragment {

    private ListView list;

    @Override
    public void onActivityCreated(@Nullable Bundle savedInstanceState) {
        super.onActivityCreated(savedInstanceState);
    }

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        getActivity().setTitle("Track your meals and workouts");
    }

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_meal, container, false);

        list = (ListView) v.findViewById(R.id.listView);

        list.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                JSONObject jo = (JSONObject) parent.getItemAtPosition(position);

                try {
                    showMealDialog(getContext(), jo.getString("timestamp"), jo.getString("description"), jo.getString("calories"));
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
                    byte[] data = Fandroid.getMeals();
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
                        return R.layout.fragment_meal_item;
                    }

                    @Override
                    public void bind(JSONObject jo, int position, View convertView, ViewGroup parent) {
                        final TextView date = (TextView) convertView.findViewById(R.id.txtDate);
                        final TextView desc = (TextView) convertView.findViewById(R.id.txtDescription);
                        final TextView cals = (TextView) convertView.findViewById(R.id.txtCalories);
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
                            cals.setText(jo.getString("calories"));

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

    public void showMealDialog(final Context context, final String id, final String desc, final String calories) {
        String negativeText = "Cancel";
        if (!id.isEmpty()) {
            negativeText = "Delete";
        }

        final String neg = negativeText;

        MaterialDialog md = new MaterialDialog.Builder(context)
                .title("Track a meal or workout")
                .customView(R.layout.dialog_addmeal, true)
                .positiveText("OK")
                .negativeText(negativeText)
                .onNegative(new MaterialDialog.SingleButtonCallback() {
                    @Override
                    public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                        if (neg.equals("Delete")) {

                            Task.callInBackground(new Callable<Object>() {
                                @Override
                                public Object call() throws Exception {
                                    Fandroid.deleteMeal(id);
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
                                EditText caloriesTxt = (EditText) d.getCustomView().findViewById(R.id.txtCalories);
                                Spinner caloricUnitList = (Spinner) d.getCustomView().findViewById(R.id.listCaloricUnit);

                                String description = descTxt.getText().toString();
                                String calories = caloriesTxt.getText().toString();
                                String unit = caloricUnitList.getSelectedItem().toString();

                                if (id.isEmpty()) {
                                    Fandroid.addNewMeal(description, calories, unit);
                                } else {
                                    Fandroid.updateMeal(id, description, calories);
                                }

                                return null;
                            }
                        }).continueWith(new Continuation<Object, Object>() {
                            @Override
                            public Object then(Task<Object> task) throws Exception {
                                Exception e = task.getError();
                                if ( e != null) {
                                    MessageUtil.showError(context, e.getMessage()).onAny(new MaterialDialog.SingleButtonCallback() {
                                        @Override
                                        public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                                            showMealDialog(context, id, desc, calories);
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
        EditText caloriesTxt = (EditText) d.findViewById(R.id.txtCalories);

        final Spinner listCaloricUnit = (Spinner) d.findViewById(R.id.listCaloricUnit);

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
                listCaloricUnit.setAdapter(units);
                return null;
            }
        }, Task.UI_THREAD_EXECUTOR);

        descTxt.setText(desc);
        caloriesTxt.setText(calories);
    }

}
