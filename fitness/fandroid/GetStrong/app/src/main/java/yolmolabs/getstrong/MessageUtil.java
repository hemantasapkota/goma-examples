package yolmolabs.getstrong;

import android.content.Context;

import com.afollestad.materialdialogs.MaterialDialog;

/**
 * Created by hemantasapkota on 27/06/16.
 */
public class MessageUtil {

    public static MaterialDialog.Builder showError(Context context, String message) {
        return new MaterialDialog.Builder(context).title("Error").content(message).positiveText("OK");
    }

}
