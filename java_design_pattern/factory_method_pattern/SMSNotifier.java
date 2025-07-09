package java_design_pattern.factory_method_pattern;

public class SMSNotifier implements Notifier{
    public void send(String message) {
        System.out.println("Sending SMS notification: " + message);
    }
}
