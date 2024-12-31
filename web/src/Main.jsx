// メインページ

import { useNavigate } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";

const Main = () => {
  const navigate = useNavigate();

  // shout ボタン押下時に実行される，フォームの提出処理をする関数 shout
  const shout = () => {
    navigate("/timeline");
  };

  return (
    <Container>
      {/* ヘッダー画像 */}
      <img
        src={header}
        alt="the yamabiko's header"
        className="d-block mt-5 mx-auto w-50"
      />

      <Form>
        <div className="d-flex gap-2 mt-3">
          {/* テキストボックス voice */}
          <Form.Group
            controlId="voice"
            className="align-self-center flex-grow-1"
          >
            <Form.Control placeholder="どんな話題がある～？" />
          </Form.Group>

          {/* フォームの送信ボタン shout  */}
          <Button type="submit" onClick={shout}>
            <img
              src={megaphone}
              alt="shout (generally means submit, search) icon"
              style={{ width: "30px" }}
            />
          </Button>
        </div>

        {/* `Demo mode results` のチェックボックス */}
        <Form.Group controlId="demoMode" className="d-flex flex-row-reverse">
          <Form.Check type="checkbox" label="Demo mode results" />
        </Form.Group>
      </Form>
    </Container>
  );
};

export default Main;
