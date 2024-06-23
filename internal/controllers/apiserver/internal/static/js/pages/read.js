const app = Vue.createApp({
  setup() {
    const bookID = new URLSearchParams(window.location.search).get("book");

    const appState = Vue.reactive({
      book: {},
      activePage: {},
      bookError: "",
      pageNumber:
        parseInt(new URLSearchParams(window.location.search).get("page")) || 1,
    });

    async function getBook() {
      axios
        .post("/api/book/details", { id: bookID })
        .then(function (response) {
          let data = response.data;
          appState.book = data;
          appState.bookError = "";
          appState.activePage = data.pages[appState.pageNumber - 1] || {};
        })
        .catch(function (error) {
          console.log(error);
          appState.bookError = error.toString();
        });
    }

    function goPage(pageNumberNew) {
      appState.activePage = appState.book.pages[pageNumberNew - 1];
      appState.pageNumber = pageNumberNew;
    }

    function prevPage() {
      if (appState.pageNumber == 1) return;

      goPage(appState.pageNumber - 1);
    }

    function nextPage() {
      if (appState.pageNumber == appState.book.pages.length) return;

      goPage(appState.pageNumber + 1);
    }

    function goGo(event) {
      const pos = document.getElementById("main-image").getBoundingClientRect();
      const dx = (event.pageX - pos.left) / (pos.right - pos.left);
      if (dx < 0.3) {
        prevPage();
      } else {
        nextPage();
      }
    }

    Vue.onBeforeMount(() => {
      getBook();

      window.addEventListener("keydown", function (event) {
        if (event.keyCode === 37) prevPage();
        if (event.keyCode === 39) nextPage();
      });
    });

    return {
      appState,
      getBook,
      goGo,
      prevPage,
      nextPage,
      goPage,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
