class Rating {
  constructor(node) {
    this.render = this.render.bind(this);
    this.updateRating = this.updateRating.bind(this);

    document.createElement("span").getAttribute("");

    this.node = node;

    this.bookID = parseInt(this.node.getAttribute("book"));
    this.rating = parseInt(this.node.getAttribute("rating") || "0");
    this.pageNumber = parseInt(this.node.getAttribute("page") || "0");

    this.render();
  }
  async updateRating(rating) {
    try {
      let response = await fetch("/api/rate", {
        method: "POST",
        body: JSON.stringify({
          id: this.bookID,
          page: this.pageNumber || 0,
          rating: rating,
        }),
      });

      if (!response.ok) {
        throw new Error(await response.text());
      }

      this.rating = rating;
      this.render();
    } catch (e) {
      console.log(e);
    }
  }
  render() {
    this.node.innerHTML = "";

    for (let index = 1; index < 6; index++) {
      let node = document.createElement("span");
      node.className = "rating-select";
      if (this.rating >= index) {
        node.setAttribute("rating", this.rating);
      }
      node.onclick = () => this.updateRating(index);
      this.node.appendChild(node);
    }
  }
}

function refreshRatings() {
  window.dispatchEvent(new Event("app-refresh-ratings"));
}

window.addEventListener("app-refresh-ratings", function () {
  document.querySelectorAll("span.rating[unprocessed]").forEach((node) => {
    new Rating(node);
    node.removeAttribute("unprocessed");
  });
});
