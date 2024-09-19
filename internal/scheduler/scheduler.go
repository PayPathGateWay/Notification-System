// internal/scheduler.go
package scheduler

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"go-notification-system/config"
//	"log"
//	"time"
//
//	"github.com/go-redis/redis/v8"
//)
//
//type Scheduler struct {
//	redisClient    *redis.Client
//	deliveryMethod DeliveryMethod
//	tmplEngine     *TemplateEngine
//	cfg            *config.Config
//}
//
//func NewScheduler(cfg *config.Config) *Scheduler {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     cfg.RedisAddr,
//		Password: cfg.RedisPassword,
//		DB:       cfg.RedisDB,
//	})
//
//	// Example template (could be dynamic based on req.Template)
//	templateContent := "Hello, {{.name}}! This is your notification."
//
//	tmplEngine, err := NewTemplateEngine(templateContent)
//	if err != nil {
//		log.Fatalf("Failed to create template engine: %v", err)
//	}
//
//	return &Scheduler{
//		redisClient:    rdb,
//		deliveryMethod: NewEmailDelivery(),
//		tmplEngine:     tmplEngine,
//		cfg:            cfg,
//	}
//}
//
//func (s *Scheduler) StartScheduler() {
//	fmt.Println("Starting scheduler...")
//	ticker := time.NewTicker(2 * time.Second) // Check every 2 seconds
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			s.processNotifications()
//		}
//	}
//}
//
//func (s *Scheduler) processNotifications() {
//	ctx := context.Background()
//
//	// Fetch the highest priority notification (lowest score)
//	// ZPOPMIN is atomic and pops the smallest element
//	notifications, err := s.redisClient.ZPopMin(ctx, "notifications", 1).Result()
//	if err != nil && err != redis.Nil {
//		log.Printf("Error fetching notifications: %v", err)
//		return
//	}
//
//	if len(notifications) == 0 {
//		// No notifications to process
//		return
//	}
//
//	for _, notif := range notifications {
//		var req NotificationRequest
//		if err := json.Unmarshal([]byte(notif.Member.(string)), &req); err != nil {
//			log.Printf("Error unmarshalling notification: %v", err)
//			continue
//		}
//
//		// Render the template
//		message, err := s.tmplEngine.Render(req.Data)
//		if err != nil {
//			log.Printf("Error rendering template: %v", err)
//			continue
//		}
//
//		// Send the notification
//		if err := s.deliveryMethod.Send(req.Recipient, message); err != nil {
//			log.Printf("Error sending notification: %v", err)
//			// Optionally, you can retry or re-enqueue the notification
//			continue
//		}
//
//		log.Printf("Notification sent to %s", req.Recipient)
//	}
//}
