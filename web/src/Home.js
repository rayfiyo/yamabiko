import { useNavigate } from "react-router-dom";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";

const Home = () => {
  const navigate = useNavigate();

  const shout = () => {
    navigate("/timeline");
  };

  return (
    <Container>
      <img
        src={header}
        alt="the yamabiko's header"
        className="d-block mt-5 mx-auto w-50"
      />
      <Form>
        <div className="d-flex gap-2 mt-3">
          <Form.Group
            controlId="shout"
            className="align-self-center flex-grow-1"
          >
            <Form.Control placeholder="どんな話題がある～？" />
          </Form.Group>

          <Button type="submit" onClick={shout}>
            <img
              src={megaphone}
              alt="shout (generally means submit) icon"
              style={{ width: "30px" }}
            />
          </Button>
        </div>

        <Form.Group controlId="demoMode" className="d-flex flex-row-reverse">
          <Form.Check type="checkbox" label="Demo mode results" />
        </Form.Group>
      </Form>
    </Container>
  );
};

export default Home;
