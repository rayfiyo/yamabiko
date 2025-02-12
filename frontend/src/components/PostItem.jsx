// src/components/PostItem.jsx

import React from "react";
import Button from "react-bootstrap/Button";
import ListGroup from "react-bootstrap/ListGroup";

import up from "../images/icons/up.svg";
import down from "../images/icons/down.svg";

const PostItem = ({ userName, content, userIcon }) => {
  return (
    <ListGroup.Item
      as="li"
      variant="light"
      className="d-flex border rounded-3 justify-content-between align-items-start my-2"
    >
      <img
        src={userIcon}
        alt={userName}
        style={{ width: "3em" }}
        className="rounded-circle flex-shrink-0 my-1 mx-2"
      />

      <div className="me-auto">
        <h6 className="my-1">{userName}</h6>
        <p className="my-1">{content}</p>

        <Button variant="none" className="me-2">
          <img
            src={up}
            alt="up icon"
            className="mx-1"
            style={{ width: "1em" }}
          />
          14
        </Button>

        <Button variant="none" className="me-2">
          <img
            src={down}
            alt="down icon"
            className="mx-1"
            style={{ width: "1em" }}
          />
          14
        </Button>
      </div>
    </ListGroup.Item>
  );
};

export default PostItem;
