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

    async function deleteBook() {
      if (!confirm("Удаление книги необратимо, продолжить?")) {
        return;
      }

      axios
        .post("/api/book/delete", { id: bookID })
        .then(function (response) {
          alert("OK");
        })
        .catch(function (error) {
          console.log(error);
          alert(error);
        });
    }

    async function verifyBook() {
      axios
        .post("/api/book/verify", { id: bookID })
        .then(function (response) {
          alert("OK");
        })
        .catch(function (error) {
          console.log(error);
          alert(error);
        });
    }

    Vue.onBeforeMount(() => {
      getBook();
    });

    return {
      appState,
      getBook,
      deleteBook,
      verifyBook,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
