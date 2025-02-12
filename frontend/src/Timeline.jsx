// src/Timeline.jsx

import React, { useEffect, useState } from "react";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import ListGroup from "react-bootstrap/ListGroup";
import PostItem from "./components/PostItem";
import { fetchHistory } from "./services/api";

const Timeline = () => {
  // 1件分のデータを格納するステート
  const [voice, setVoice] = useState("");
  const [responses, setResponses] = useState([]);

  useEffect(() => {
    (async () => {
      try {
        const data = await fetchHistory();

        // 例: [{ ID:1, Voice:"...", Response1:"...", ...}]
        if (data.length > 0) {
          // 例えば先頭(0番目)を使う/または最後(data[data.length - 1])を使う、など運用次第
          const latest = data[0];
          setVoice(latest.Voice);

          // latest.Response1～Response6 をまとめる
          const arr = [
            latest.Response1,
            latest.Response2,
            latest.Response3,
            latest.Response4,
            latest.Response5,
            latest.Response6,
          ];
          setResponses(arr);
        }
      } catch (error) {
        console.error("Failed to fetch history:", error);
      }
    })();
  }, []);

  return (
    <Container>
      {/* Voice を表示するためのテキストボックス (readOnly) */}
      <Form.Control
        value={voice}
        aria-label="The topic you shouted out (you want to research)"
        className="mt-5 mb-3 mx-auto"
        type="text"
        readOnly
      />

      <ListGroup>
        {/* 6つの Response を 6個のカードとして表示 */}
        {responses.map((resp, idx) => (
          <PostItem
            key={idx}
            userName={`ユーザー #${idx + 1}`}
            content={resp}
            userIcon="https://github.com/twbs.png"
          />
        ))}
      </ListGroup>
    </Container>
  );
};

export default Timeline;
