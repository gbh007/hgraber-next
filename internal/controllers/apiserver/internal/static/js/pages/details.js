const app = Vue.createApp({
  setup() {
    const bookID = new URLSearchParams(window.location.search).get("book");

    const appState = Vue.reactive({
      book: {},
      bookError: "",
    });

    async function getBook() {
      axios
        .post("/api/book/details", { id: bookID })
        .then(function (response) {
          let data = response.data;
          appState.book = data;
          appState.bookError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.bookError = error.toString();
        });
    }

    Vue.onBeforeMount(() => {
      getBook();
    });

    return {
      appState,
      getBook,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
