package java_design_pattern.factory_method_pattern;

public class EmailNotifier implements Notifier {
    @Override
    public void send(String message) {
        System.out.println("Sending email notification: " + message);
    }    
}
