package yolmolabs.getstrong.fragments;

import android.content.Context;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v4.app.Fragment;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.DatePicker;
import android.widget.EditText;
import android.widget.ListView;
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
import yolmolabs.getstrong.R;

/**
 * Created by hemantasapkota on 29/1/17.
 */

public class WeightLogFragment extends Fragment {

    private ListView list;

    @Override
    public void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        getActivity().setTitle("Log weight");
    }

    @Nullable
    @Override
    public View onCreateView(LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_weightlog, container, false);

        list = (ListView) v.findViewById(R.id.listView);

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
                     byte[] data = Fandroid.getWeightLog();
                     String json = new String(data, "UTF-8");
                     jarr = new JSONArray(json);
                     for(int i = 0; i < jarr.length(); i++) {
                         jo.add(jarr.getJSONObject(i));
                     }
                 } catch (Exception e) {
                     e.printStackTrace();
                 }

                 //
                 JSONAdapter ja = new JSONAdapter(getContext(), jo) {
                     @Override
                     public int getLayoutID() {
                         return R.layout.fragment_weightlog_item;
                     }

                     @Override
                     public void bind(JSONObject jo, int position, View convertView, ViewGroup parent) {
                         final TextView weightTxt = (TextView) convertView.findViewById(R.id.txtWeight);

                         Log.d("YOLMO", "Are we here ?");

                         try {
                             weightTxt.setText(jo.getString("weight"));
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

    public void showInputDialog(final Context context, final String id) {
        String negativeText = "Cancel";
        if (!id.isEmpty()) {
            negativeText = "Delete";
        }

        final String neg = negativeText;

        MaterialDialog md = new MaterialDialog.Builder(context)
                .title("Log weight")
                .customView(R.layout.dialog_logweight, true)
                .positiveText("OK")
                .negativeText(negativeText)
                .onNegative(new MaterialDialog.SingleButtonCallback() {
                    @Override
                    public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                        if (neg.equals("Delete")) {

                            Task.callInBackground(new Callable<Object>() {
                                @Override
                                public Object call() throws Exception {
//                                    Fandroid.logWeight(10);
                                    return null;
                                }
                            }).continueWith(new Continuation<Object, Object>() {
                                @Override
                                public Object then(Task<Object> task) throws Exception {
//                                    loadData();
                                    return null;
                                }
                            }, Task.UI_THREAD_EXECUTOR);

                        }
                    }
                }).onPositive(new MaterialDialog.SingleButtonCallback() {
                    @Override
                    public void onClick(@NonNull MaterialDialog dialog, @NonNull DialogAction which) {
                        final MaterialDialog d = dialog;
                        final EditText weightTxt = (EditText) d.getCustomView().findViewById(R.id.txtWeight);
                        final DatePicker dp = (DatePicker) d.getCustomView().findViewById(R.id.datePicker);

                        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssZ");

                        final String timestamp = sdf.format(new Date(dp.getYear(), dp.getMonth(), dp.getDayOfMonth()));
                        final String weight = weightTxt.getText().toString();

                        Task.callInBackground(new Callable<Object>() {
                                @Override
                                public Object call() throws Exception {
                                    Fandroid.logWeight(timestamp, weight);
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
                }).show();

    }
}
