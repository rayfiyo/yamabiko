// src/Root.jsx

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import Alert from "react-bootstrap/Alert";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";
import { shoutVoice } from "./services/api";

const Root = () => {
  const [demoMode, setDemoMode] = useState(true);
  const [error, setError] = useState("");
  const [voice, setVoice] = useState("");
  const navigate = useNavigate();

  // フォーム送信イベントハンドラ
  const handleShout = async (e) => {
    e.preventDefault();

    // バリデーション
    if (!voice.trim()) {
      setError("テキストボックスを空にすることはできません");
      return;
    }

    // エラーリセット
    setError("");

    try {
      // shoutVoice() を呼び出して実際の API 通信を行う
      await shoutVoice({ voice, demoMode });
      // 成功したらタイムラインに移動
      navigate("/timeline");
    } catch (err) {
      // API 側やネットワークエラーはここでキャッチ
      setError(
        err.message || "予期せぬエラー: しばらくしてからもう一度お試しください"
      );
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

      <Form onSubmit={handleShout} className="mt-3">
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

          {/* フォームの送信ボタン (shout) */}
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

export default Root;
