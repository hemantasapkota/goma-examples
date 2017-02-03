package yolmolabs.getstrong;

import android.content.Context;
import android.widget.DatePicker;

import com.afollestad.materialdialogs.MaterialDialog;

import java.util.Calendar;

/**
 * Created by hemantasapkota on 27/06/16.
 */
public class AppUtil {

    public static MaterialDialog.Builder showError(Context context, String message) {
        return new MaterialDialog.Builder(context).title("Error").content(message).positiveText("OK");
    }

    public static java.util.Date getDateFromDatePicker(DatePicker datePicker){
        int day = datePicker.getDayOfMonth();
        int month = datePicker.getMonth();
        int year =  datePicker.getYear();

        Calendar calendar = Calendar.getInstance();
        calendar.set(year, month, day);

        return calendar.getTime();
    }

}
