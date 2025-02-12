// src/Timeline.jsx

import React from "react";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import ListGroup from "react-bootstrap/ListGroup";
import PostItem from "./components/PostItem";

const Timeline = () => {
  // 実際には API 等から取得したデータを入れる想定
  const mockPosts = [
    {
      id: 1,
      userName: "hoge hoge 男",
      content: "React is the library for web and native user interfaces...",
      userIcon: "https://github.com/twbs.png",
    },
    {
      id: 2,
      userName: "fuga fuga 女",
      content: "Second post sample",
      userIcon: "https://github.com/twbs.png",
    },
  ];

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
        {mockPosts.map((post) => (
          <PostItem
            key={post.id}
            userName={post.userName}
            content={post.content}
            userIcon={post.userIcon}
          />
        ))}
      </ListGroup>
    </Container>
  );
};

export default Timeline;
