package java_design_pattern.strategy_pattern;

public class NotificationService {
    private Notifier notifier;

    public NotificationService(Notifier notifier) {
        this.notifier = notifier;
    }

    public void sendNotification(String message) {
        notifier.send(message);
    }

    public void setNotifier(Notifier notifier) {
        this.notifier = notifier;
    }
}
