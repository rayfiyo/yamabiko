import React from "react";
import { Container, Form, Button } from "react-bootstrap";
import "./App.css";
// import SocialTimelineSearchResult from "./SocialTimelineSearchResult";
// <SocialTimelineSearchResult />

import header from "./images/yamabiko-header.png";
import megaphone from "./images/icons/megaphone.svg";

const App = () => {
  return (
    <Container fluid>
      <img
        src={header}
        alt="the yamabiko's header"
        className="img-fluid d-block mx-auto mt-5 w-50"
      />
      <Form>
        <div className="d-flex gap-1 mt-3">
          <Form.Group
            controlId="shout"
            className="flex-grow-1 align-self-center"
          >
            <Form.Control placeholder="どんな話題がある～？" />
          </Form.Group>

          <Button type="submit">
            <img
              src={megaphone}
              alt="shout (generally means submit) icon"
              className="img-fluid"
              style={{ width: "30px" }}
            />
          </Button>
        </div>

        <Form.Group controlId="demoMode" className="d-flex flex-row-reverse">
          <Form.Check type="checkbox" label="Demo mode results" />
        </Form.Group>
      </Form>
      ( // SocialTimelineSearchResult )
    </Container>
  );
};

export default App;
