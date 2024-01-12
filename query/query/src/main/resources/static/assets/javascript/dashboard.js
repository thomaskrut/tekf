document.addEventListener('alpine:init', () => {
    Alpine.data('row', () => ({
        readyToCheckIn: true,

        clean: function (button) {
            button.disabled = true;
        },

        checkin: function (button) {
            this.postCheckin(button)
        },

        checkout: function (button) {
            this.postCheckout(button)
            this.readyToCheckIn = true
        },

        postCheckin: function (button) {
            const id = button.dataset.id
            fetch('http://localhost:8080/checkin/' + id, {
                method: 'POST',
            }).then(response => {
                if (response.ok) {
                    window.location.reload()
                } else {
                    alert('Check in failed: ' + response.status)
                }
            })
        },

        postCheckout: function (button) {
            const id = button.dataset.id
            fetch('http://localhost:8080/checkout/' + id, {
                method: 'POST',
            }).then(response => {
                if (response.ok) {
                    window.location.reload()
                } else {
                    alert('Check out failed: ' + response.status)
                }
            })
        },
    }))
})