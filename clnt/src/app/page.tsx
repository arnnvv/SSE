"use client";

import { useEffect, JSX, useState } from "react";

export default function Home(): JSX.Element {
  const [messages, setMessages] = useState<string[]>([]);

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:3001");

    eventSource.onmessage = (event: MessageEvent) => {
      setMessages((prev) => [...prev, event.data]);
    };

    eventSource.onerror = () => {
      console.error("EventSource failed.");
      eventSource.close();
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return (
    <div>
      <h1>SSE Client</h1>
      <ul>
        {messages.map((msg, index) => (
          <li key={index}>{msg}</li>
        ))}
      </ul>
    </div>
  );
}
