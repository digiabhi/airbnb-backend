export interface NotificationDTO {
  to: string;
  subject: string;
  tempelateId: string;
  params: Record<string, any>;
}
