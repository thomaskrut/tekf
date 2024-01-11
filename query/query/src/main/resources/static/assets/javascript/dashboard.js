document.addEventListener('alpine:init', () => {
    Alpine.data('row', () => ({
        readyToCheckIn: true,

        clean: function (button) {
            button.disabled = true;
        },

        checkout: function (button) {
            this.readyToCheckIn = true
        },
    }))
})