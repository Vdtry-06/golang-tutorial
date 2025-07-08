package java_design_pattern.strategy_pattern;

public class SMSNotifier implements Notifier {
    @Override
    public void send(String message) {
        System.out.println("Sending SMS with message: " + message);
    }   
}
