const deleteConfirmDialog = document.querySelector("#deleteConfirmDialog");
const deleteConfirmButton = document.querySelector(".delete-confirm-button");
const deleteCancelButton = document.querySelector(".delete-cancel-button");
const deleteForms = document.querySelectorAll(".delete-form");

let pendingDeleteForm = null;

function enableDeleteDialogArrowNavigation() {
    if (!deleteConfirmDialog) {
        return;
    }
    const buttons = [deleteCancelButton, deleteConfirmButton];

    deleteConfirmDialog.addEventListener("keydown", function (event) {
        if (
            event.key !== "ArrowLeft" &&
            event.key !== "ArrowRight" &&
            event.key !== "ArrowUp" &&
            event.key !== "ArrowDown"
        ) {
            return;
        }

        event.preventDefault();

        const step = event.key === "ArrowRight" || event.key === "ArrowDown" ? 1 : -1;
        focusElementByStep(buttons, step);
    });
}

if (deleteConfirmDialog) {
    enableDeleteDialogArrowNavigation();

    deleteForms.forEach(function (form) {
        form.addEventListener("submit", function (event) {
            event.preventDefault();
            pendingDeleteForm = form;
            deleteConfirmDialog.showModal();
            deleteCancelButton?.focus();
        });
    });

    deleteConfirmButton?.addEventListener("click", function () {
        pendingDeleteForm?.submit();
    });

    deleteCancelButton?.addEventListener("click", function () {
        pendingDeleteForm = null;
        deleteConfirmDialog.close();
    });
}
