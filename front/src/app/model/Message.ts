

export class Message {
  content: string;
  sender: string | null;
  receiver: string;

  constructor() {
    this.content = '';
    this.sender = '';
    this.receiver = '';
  }
}
