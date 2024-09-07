const app = Vue.createApp({
  setup() {
    let countOnPage = getBookOnPageCount();

    function getBookOnPageCount() {
      let data = JSON.parse(localStorage.getItem("settings")) || {};
      return data.book_on_page || 12;
    }

    const appState = Vue.reactive({
      books: [],
      bookCount: 0,
      booksError: "",
      pages: [],
      showDeleted: "except",
      showVerify: "only",
      showDownloaded: "only",
      nameFilter: "",
      agents: [],
      agentStatusError: "",
      exportAgentID: "",
      deleteAfterExport: false,
      from: "",
      to: "",
      sortDesc: true,
      sortField: "created_at",
    });

    async function renderPages(pageNumber = 1) {
      let filter = {
        count: countOnPage,
        page: pageNumber,
        verify_status: appState.showVerify,
        delete_status: appState.showDeleted,
        download_status: appState.showDownloaded,
        sort_desc: appState.sortDesc,
        sort_field: appState.sortField,
        filter: {
          name: appState.nameFilter,
        },
      };

      if (appState.from) {
        filter.from = new Date(appState.from).toJSON();
      }

      if (appState.to) {
        filter.to = new Date(appState.to).toJSON();
      }

      axios
        .post("/api/book/list", filter)
        .then(function (response) {
          let data = response.data;
          appState.books = data.books;
          appState.pages = data.pages;
          appState.bookCount = data.count || 0;
          appState.booksError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.booksError = error.toString();
        });
    }

    function agents() {
      axios
        .post("/api/agent/list", {
          can_export: true,
        })
        .then(function (response) {
          appState.agents = response.data;
          appState.agentStatusError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.agentStatusError = error.toString();
        });
    }

    function exportBooks() {
      let filter = {
        verify_status: appState.showVerify,
        delete_status: appState.showDeleted,
        download_status: appState.showDownloaded,
        sort_desc: appState.sortDesc,
        sort_field: appState.sortField,
        filter: {
          name: appState.nameFilter,
        },
      };

      if (appState.from) {
        filter.from = new Date(appState.from).toJSON();
      }

      if (appState.to) {
        filter.to = new Date(appState.to).toJSON();
      }

      axios
        .post("/api/agent/task/export", {
          book_filter: filter,
          exporter: appState.exportAgentID,
          delete_after: appState.deleteAfterExport,
        })
        .then(function () {})
        .catch(function (error) {
          console.log(error);
          alert(error.toString());
        });
    }

    Vue.onBeforeMount(() => {
      agents();
      renderPages();
    });

    return {
      appState,
      renderPages,
      agents,
      exportBooks,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
