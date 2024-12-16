import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import ListGroup from "react-bootstrap/ListGroup";

import down from "./images/icons/down.svg";
import up from "./images/icons/up.svg";

// import React, { useState } from "react";
// import { faSearch } from "@fortawesome/free-solid-svg-icons";
// import { useHistory } from "react-router-dom";

const Timeline = () => {
  const userName = "hoge hoge 男";
  const content =
    "React is the library for web and native user interfaces. Build user interfaces out of individual pieces called components written in JavaScript.";
  const posts = (
    <ListGroup.Item
      as="li"
      variant="light"
      className="d-flex border  rounded-3 justify-content-between align-items-start my-2"
    >
      <img
        src="https://github.com/twbs.png"
        alt="twbs"
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
            alt="up icon"
            className="mx-1"
            style={{ width: "1em" }}
          />
          14
        </Button>
      </div>
    </ListGroup.Item>
  );
  return (
    <Container>
      <Form.Control
        placeholder="$（話題）"
        aria-label="The topic you shouted out (you want to research)"
        className="mt-5 mb-3 mx-auto"
        type="text"
        readOnly
      />

      <ListGroup>
        {posts}
        {posts}
      </ListGroup>
    </Container>
  );
};

export default Timeline;

/*
  const [searchQuery, setSearchQuery] = useState("");
  // const [posts, setPosts] = useState([
  // 投稿データ;
  // ]);
  // const history = useHistory();

  const handleSearch = (e) => {
    setSearchQuery(e.target.value);
    // 検索機能の実装
  };

  const handleSubmit = () => {
    // 検索クエリをURLのパラメータとして送る
    // history.push(`/search?q=${encodeURIComponent(searchQuery)}`);
  };



      <div className="container my-5">
        <div className="row mb-4">
          <div className="col">
            <div className="input-group">
              <input
                type="text"
                className="form-control"
                placeholder="Search posts..."
                value={searchQuery}
                onChange={handleSearch}
              />
            </div>
          </div>
        </div>

        <div className="row">{
        // 投稿一覧の表示
        }</div>
      </div>
*/
