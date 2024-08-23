const app = Vue.createApp({
  setup() {
    let countOnPage = getBookOnPageCount();

    function getBookOnPageCount() {
      let data = JSON.parse(localStorage.getItem("settings")) || {};
      return data.book_on_page || 12;
    }

    const appState = Vue.reactive({
      books: [],
      booksError: "",
      pages: [],
      showDeleted: "except",
      showVerify: "only",
      nameFilter: "",
    });

    async function renderPages(pageNumber = 1) {
      axios
        .post("/api/book/list", {
          count: countOnPage,
          page: pageNumber,
          verify_status: appState.showVerify,
          delete_status: appState.showDeleted,
          filter: {
            name: appState.nameFilter,
          },
        })
        .then(function (response) {
          let data = response.data;
          appState.books = data.books;
          appState.pages = data.pages;
          appState.booksError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.booksError = error.toString();
        });
    }

    Vue.onBeforeMount(() => {
      renderPages();
    });

    return {
      appState,
      renderPages,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
