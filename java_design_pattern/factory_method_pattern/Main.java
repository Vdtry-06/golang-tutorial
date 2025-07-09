package java_design_pattern.factory_method_pattern;

public class Main {
    public static void main(String[] args) {
        Notifier notifier = NotifierFactory.createNotifier("Email");
        NotificationService notificationService = new NotificationService(notifier);
        notificationService.sendNotification("Hello, this is a test notification!");
    }
}