document.addEventListener("DOMContentLoaded", () => {
  const dropZone = document.getElementById("drop-zone");
  const fileInput = document.getElementById("file-input");
  const btnSelect = document.getElementById("btn-select");
  const filePreview = document.getElementById("file-preview");
  const fileNameDisplay = document.getElementById("file-name");
  const fileSizeDisplay = document.getElementById("file-size");
  const fileIconDisplay = document.getElementById("file-icon-display");
  const btnRemove = document.getElementById("btn-remove");
  const btnUpload = document.getElementById("btn-upload");
  const btnText = btnUpload.querySelector(".btn-text");
  const loader = btnUpload.querySelector(".loader");
  const successIcon = btnUpload.querySelector(".success-icon");

  let currentFile = null;

  // Trigger file selection
  btnSelect.addEventListener("click", () => {
    fileInput.click();
  });

  fileInput.addEventListener("change", (e) => {
    handleFileSelect(e.target.files[0]);
  });

  // Drag and Drop Handling
  ["dragenter", "dragover", "dragleave", "drop"].forEach((eventName) => {
    dropZone.addEventListener(eventName, preventDefaults, false);
  });

  function preventDefaults(e) {
    e.preventDefault();
    e.stopPropagation();
  }

  ["dragenter", "dragover"].forEach((eventName) => {
    dropZone.addEventListener(eventName, highlight, false);
  });

  ["dragleave", "drop"].forEach((eventName) => {
    dropZone.addEventListener(eventName, unhighlight, false);
  });

  function highlight(e) {
    dropZone.classList.add("drag-over");
  }

  function unhighlight(e) {
    dropZone.classList.remove("drag-over");
  }

  dropZone.addEventListener("drop", (e) => {
    const dt = e.dataTransfer;
    const file = dt.files[0];
    handleFileSelect(file);
  });

  // Handle selected file
  function handleFileSelect(file) {
    if (!file) return;

    currentFile = file;
    fileNameDisplay.textContent = file.name;
    fileSizeDisplay.textContent = formatFileSize(file.size);

    // Show file preview
    dropZone.classList.add("has-file");
    btnUpload.disabled = false;

    // Generate preview icon
    if (file.type.startsWith("image/")) {
      const reader = new FileReader();
      reader.onload = (e) => {
        fileIconDisplay.innerHTML = `<img src="${e.target.result}" alt="Preview">`;
      };
      reader.readAsDataURL(file);
    } else {
      fileIconDisplay.innerHTML = `
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M13 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9z"></path>
                    <polyline points="13 2 13 9 20 9"></polyline>
                </svg>
            `;
    }
  }

  // Format file size
  function formatFileSize(bytes) {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const sizes = ["Bytes", "KB", "MB", "GB", "TB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
  }

  // Remove file
  btnRemove.addEventListener("click", () => {
    currentFile = null;
    fileInput.value = "";
    dropZone.classList.remove("has-file");
    btnUpload.disabled = true;
    resetUploadButton();
  });

  // Handle Upload
  btnUpload.addEventListener("click", () => {
    if (!currentFile) return;

    // Simulate upload
    startUploadSimulation();
  });

  function startUploadSimulation() {
    btnUpload.classList.add("uploading");
    btnText.textContent = "Uploading...";
    loader.classList.remove("hidden");

    const formData = new FormData();
    formData.append("file", currentFile);

    fetch("http://localhost:8080/upload", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (response.ok) {
          finishUpload();
        } else {
          throw new Error("Upload failed");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
        btnUpload.classList.remove("uploading");
        loader.classList.add("hidden");
        btnText.textContent = "Failed";
        setTimeout(() => {
          resetUploadButton();
        }, 2000);
      });
  }

  function finishUpload() {
    btnUpload.classList.remove("uploading");
    btnUpload.classList.add("success");
    loader.classList.add("hidden");
    btnText.textContent = "Uploaded";
    successIcon.classList.remove("hidden");

    // Reset after success
    setTimeout(() => {
      console.log("Upload completed");
      resetUploadButton();
      // Optional: Auto clear file after success
      // btnRemove.click();
    }, 2500);
  }

  function resetUploadButton() {
    btnUpload.classList.remove("uploading", "success");
    btnText.textContent = "Upload File";
    loader.classList.add("hidden");
    successIcon.classList.add("hidden");
  }
});
