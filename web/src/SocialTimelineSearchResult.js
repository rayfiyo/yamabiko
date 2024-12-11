import React, { useState } from "react";
// import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
// import { faSearch } from "@fortawesome/free-solid-svg-icons";
// import { useHistory } from "react-router-dom";

const SocialTimelineSearchResult = () => {
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

  return (
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
            <div className="input-group-append">
              <button
                className="btn btn-primary"
                type="button"
                onClick={handleSubmit}
              >
                "// FontAwesomeIcon"
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="row">{/* 投稿一覧の表示 */}</div>
    </div>
  );
};

export default SocialTimelineSearchResult;
// <FontAwesomeIcon icon={faSearch} />
