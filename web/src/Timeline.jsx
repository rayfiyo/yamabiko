import Badge from "react-bootstrap/Badge";
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import ListGroup from "react-bootstrap/ListGroup";
// import React, { useState } from "react";
// import { faSearch } from "@fortawesome/free-solid-svg-icons";
// import { useHistory } from "react-router-dom";

const Timeline = () => {
  return (
    <Container>
      <Form.Control
        placeholder="$（話題）"
        aria-label="The topic you shouted out (you want to research)"
        className="my-5 mx-auto"
        type="text"
        readOnly
      />

      <ListGroup.Item
        as="li"
        className="d-flex justify-content-between align-items-start"
      >
        <img
          src="https://github.com/twbs.png"
          alt="twbs"
          style={{ width: "32px" }}
          className="my-0 rounded-circle flex-shrink-0"
        />
        <div className="ms-2 me-auto">
          <h6>List group item heading</h6>
          <p class="opacity-75">
            Some placeholder content in a paragraph. Some placeholder content in
            a paragraph Some placeholder content in a paragraph Some placeholder
            content in a paragraph Some placeholder content in a paragraph Some
            placeholder content in a paragraph
          </p>
        </div>
        <Badge bg="primary" pill>
          14
        </Badge>
      </ListGroup.Item>
    </Container>
  );
};

export default Timeline;

/*
    <a href="#" class="list-group-item list-group-item-action d-flex gap-3 py-3" aria-current="true">
      <img src="https://github.com/twbs.png" alt="twbs" width="32" height="32" class="rounded-circle flex-shrink-0">
      <div class="d-flex gap-2 w-100 justify-content-between">
        <div>
          <h6 class="mb-0">List group item heading</h6>
          <p class="mb-0 opacity-75">Some placeholder content in a paragraph.</p>
        </div>
        <small class="opacity-50 text-nowrap">now</small>
      </div>
    </a>
*/

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
