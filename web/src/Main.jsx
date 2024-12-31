// メインページ

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import Alert from "react-bootstrap/Alert";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";

const Main = () => {
  const [demoMode, setDemoMode] = useState(true);
  const [error, setError] = useState("");
  const [voice, setVoice] = useState("");
  const navigate = useNavigate();

  // shout ボタン押下時に実行される，フォームの提出処理をする関数 shout
  const shout = async (e) => {
    // フォームのデフォルトの動作（ページリロード）をキャンセル
    e.preventDefault();

    // バリデーション
    if (!voice.trim()) {
      setError("テキストボックスを空にすることはできません");
      console.error("Textbox cannot be empty.");
      return;
    }

    // エラーリセット
    setError("");

    // エラーハンドリング付きで，voice をエンドポイントに shout
    try {
      // 成功時: voice を shout
      const response = await fetch("/api/shout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ voice, demoMode }),
      });

      // エンドポイントに POST した後のエラーハンドリング
      if (!response.ok) {
        const errorData = await response.json();
        setError(errorData.message || "shout に失敗");
        console.error(errorData.message || "Failed to shout.");
        return;
      }

      // ページの遷移
      navigate("/timeline");
    } catch (error) {
      // エラーハンドリング
      setError("予期せぬエラー: しばらくしてからもう一度お試しください");
      console.error("An unexpected error: Please try again later.");
    }
  };

  return (
    <Container>
      {/* ヘッダー画像 */}
      <img
        src={header}
        alt="The Yamabiko header"
        className="d-block mt-5 mx-auto w-50"
      />

      <Form onSubmit={shout} className="mt-3">
        {error && (
          <Alert variant="danger" onClose={() => setError("")} dismissible>
            {error}
          </Alert>
        )}

        <div className="d-flex gap-2">
          {/* テキストボックス voice */}
          <Form.Group
            controlId="voice"
            className="align-self-center flex-grow-1"
          >
            <Form.Control
              placeholder="どんな話題がある～？"
              value={voice}
              onChange={(e) => setVoice(e.target.value)}
              aria-label="Voice input box (Use to shout)"
            />
          </Form.Group>

          {/* フォームの送信ボタン shout  */}
          <Button type="submit" aria-label="Shout button">
            <img src={megaphone} alt="Shout icon" style={{ width: "30px" }} />
          </Button>
        </div>

        {/* `Demo mode results` のチェックボックス */}
        <Form.Group
          controlId="demoMode"
          className="d-flex flex-row-reverse mt-2"
        >
          <Form.Check
            disabled
            type="checkbox"
            checked={demoMode}
            onChange={(e) => setDemoMode(e.target.checked)}
            label="Demo mode results"
          />
        </Form.Group>
      </Form>
    </Container>
  );
};

export default Main;
