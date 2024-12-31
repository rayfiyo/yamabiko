// メインページ

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";

const Main = () => {
  const [voice, setVoice] = useState("");
  const navigate = useNavigate();

  // shout ボタン押下時に実行される，フォームの提出処理をする関数 shout
  const shout = async (e) => {
    // フォームのデフォルトの動作（ページリロード）をキャンセル
    e.preventDefault();

    // バリデーション
    if (!voice.trim()) {
      console.error("Message cannot be empty.");
      return;
    }

    // エラーハンドリング付きで，voice をエンドポイントに shout
    try {
      // 成功時: voice を shout
      const response = await fetch("/api/shout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ voice }),
      });

      // エンドポイントに POST した後のエラーハンドリング
      if (!response.ok) {
        const errorData = await response.json();
        console.error(errorData.message || "Failed to shout.");
        return;
      }

      // ページの遷移
      navigate("/timeline");
    } catch (error) {
      // エラーハンドリング
      console.error("An unexpected error occurred:", error);
    }
  };

  return (
    <Container>
      {/* ヘッダー画像 */}
      <img
        src={header}
        alt="the yamabiko's header"
        className="d-block mt-5 mx-auto w-50"
      />

      <Form onSubmit={shout}>
        <div className="d-flex gap-2 mt-3">
          {/* テキストボックス voice */}
          <Form.Group
            controlId="voice"
            className="align-self-center flex-grow-1"
          >
            <Form.Control
              placeholder="どんな話題がある～？"
              value={voice}
              onChange={(e) => setVoice(e.target.value)}
            />
          </Form.Group>

          {/* フォームの送信ボタン shout  */}
          <Button type="submit">
            <img
              src={megaphone}
              alt="shout (generally means submit, search) icon"
              style={{ width: "30px" }}
            />
          </Button>
        </div>

        {/* `Demo mode results` のチェックボックス */}
        <Form.Group controlId="demoMode" className="d-flex flex-row-reverse">
          <Form.Check
            disabled
            type="checkbox"
            checked="true"
            label="Demo mode results"
          />
        </Form.Group>
      </Form>
    </Container>
  );
};

export default Main;
