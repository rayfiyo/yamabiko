import React from "react";
import { Container, Row, Col, Form, Button } from "react-bootstrap";
import SocialTimelineSearchResult from "./SocialTimelineSearchResult";

const App = () => {
  return (
    <Container fluid>
      <Row className="justify-content-center mt-5">
        <Col xs={12} md={8} lg={6}>
          <img
            src="images/yamabiko-header.png"
            alt="Mountain"
            className="img-fluid"
          />
          <Form className="mt-4">
            <Form.Group controlId="question">
              <Form.Label>どんな話題がありますか?</Form.Label>
              <Form.Control as="textarea" rows={3} />
            </Form.Group>
            <div className="d-grid gap-2">
              <Button variant="primary" type="submit">
                <img
                  src="public/images/trumpet_icon_24349.svg"
                  alt="Icon"
                  className="me-2"
                />
                Submit
              </Button>
            </div>
          </Form>
        </Col>
      </Row>
      <SocialTimelineSearchResult />
    </Container>
  );
};

export default App;
