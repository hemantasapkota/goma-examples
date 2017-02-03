package yolmolabs.getstrong;

import android.app.Application;

import go.fandroid.Fandroid;
import uk.co.chrisjenx.calligraphy.CalligraphyConfig;

/**
 * Created by hemantasapkota on 27/06/16.
 */
public class FApplication extends Application {

    @Override
    public void onCreate() {
        super.onCreate();
        try {
            Fandroid.init(getFilesDir().toString());
        } catch (Exception e) {
            e.printStackTrace();
        }

        CalligraphyConfig.initDefault(new CalligraphyConfig.Builder()
                .setDefaultFontPath("fonts/Roboto-RobotoRegular.ttf")
                .setFontAttrId(R.attr.fontPath)
                .build()
        );

    }

    @Override
    public void onTerminate() {
        super.onTerminate();
    }
}
