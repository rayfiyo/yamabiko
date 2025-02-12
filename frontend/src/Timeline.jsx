// src/Timeline.jsx
import React, { useEffect, useState } from "react";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import ListGroup from "react-bootstrap/ListGroup";
import PostItem from "./components/PostItem";

// 先ほど追加した fetchHistory をインポート
import { fetchHistory } from "./services/api";

const Timeline = () => {
  const [histories, setHistories] = useState([]);

  useEffect(() => {
    // コンポーネントがマウントされたときに /api/history を取得
    (async () => {
      try {
        const data = await fetchHistory();
        // data は配列
        // 例: [{ ID: 1, Voice: "…", Response1: "…", ..., CreatedAt: "…" }, ...]
        setHistories(data);
      } catch (error) {
        console.error("Failed to fetch history:", error);
      }
    })();
  }, []);

  return (
    <Container>
      {/* 送信した話題を表示する想定のテキストボックス。要件によって活用方法を調整。 */}
      <Form.Control
        placeholder="$（話題）"
        aria-label="The topic you shouted out (you want to research)"
        className="mt-5 mb-3 mx-auto"
        type="text"
        readOnly
      />

      <ListGroup>
        {histories.map((h) => (
          <PostItem
            key={h.ID}
            userName={`(ID:${h.ID})`} // 適当に
            content={`Voice: ${h.Voice}\n1) ${h.Response1}`}
            userIcon="https://github.com/twbs.png"
          />
        ))}
      </ListGroup>
    </Container>
  );
};

export default Timeline;
