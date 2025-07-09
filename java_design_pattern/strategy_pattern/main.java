package java_design_pattern.strategy_pattern;

public class Main {
    public static void main(String[] args) {
        NotificationService notificationService = new NotificationService(new EmailNotifier());
        notificationService.sendNotification("Hello via Email!");
        notificationService.setNotifier(new SMSNotifier());
        notificationService.sendNotification("Hello via SMS!");
    }
}
