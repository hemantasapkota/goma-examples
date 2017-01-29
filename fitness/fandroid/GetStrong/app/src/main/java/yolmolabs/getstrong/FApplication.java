package yolmolabs.getstrong;

import android.app.Application;

import go.fandroid.Fandroid;

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
    }

    @Override
    public void onTerminate() {
        super.onTerminate();
    }
}
