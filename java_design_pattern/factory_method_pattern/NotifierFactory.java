package java_design_pattern.factory_method_pattern;

public class NotifierFactory {
    public static Notifier createNotifier(String type) {
        switch (type) {
            case "SMS":
                return new SMSNotifier();
            case "Email":
                return new EmailNotifier();
            default:
                throw new IllegalArgumentException("Unknown notifier type: " + type);
        }
    }
}
