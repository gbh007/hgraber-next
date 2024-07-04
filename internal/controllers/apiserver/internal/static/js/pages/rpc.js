const app = Vue.createApp({
  setup() {
    const appState = Vue.reactive({
      deduplicateFiles: null,
      deduplicateFilesError: "",
      removeDetachedFiles: null,
      removeDetachedFilesError: "",
      removeMismatchFiles: null,
      removeMismatchFilesError: "",
    });

    function deduplicateFiles() {
      axios
        .post("/api/system/rpc/deduplicate/files")
        .then(function (response) {
          appState.deduplicateFiles = response.data;
          appState.deduplicateFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.deduplicateFilesError = error.toString();
        });
    }

    function removeDetachedFiles() {
      axios
        .post("/api/system/rpc/remove/detached-files")
        .then(function (response) {
          appState.removeDetachedFiles = response.data;
          appState.removeDetachedFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.removeDetachedFilesError = error.toString();
        });
    }

    function removeMismatchFiles() {
      axios
        .post("/api/system/rpc/remove/mismatch-files")
        .then(function (response) {
          appState.removeMismatchFiles = response.data;
          appState.removeMismatchFilesError = "";
        })
        .catch(function (error) {
          console.log(error);
          appState.removeMismatchFilesError = error.toString();
        });
    }

    return {
      appState,
      deduplicateFiles,
      removeDetachedFiles,
      removeMismatchFiles,
    };
  },
});

window.addEventListener("load", async function () {
  app.mount("#app");
});
