package java_design_pattern.strategy_pattern;

public class main {
    public static void main(String[] args) {
        NotificationService notificationService = new NotificationService(new EmailNotifier());
        notificationService.SendNotification("Hello via Email!");
        notificationService.setNotifier(new SMSNotifier());
        notificationService.SendNotification("Hello via SMS!");
    }
}
