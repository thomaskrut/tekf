const calendar = {
    bookingId: '',
    name: '',
    guests: 1,
    unitId: null,
    fromDate: null,
    toDate: null,
    fromDateParts: null,
    selectedCells: [],
    highlitCells: [],
    originalColor: '',
    lastCell: null,
    makingSelection: false,
    selectionMade: false,
    bookingSelected: false,

    clickDate: function (td) {

        if (this.selectionMade || this.bookingSelected) {
            this.clear()
        }

        if (td.classList.contains('booked') && !this.makingSelection) {
            this.bookingSelected = true
            this.showBooking(td)
            this.highlightCells(td)
            return
        }

        if (!td.classList.contains('booked') && !this.makingSelection) {
            this.makingSelection = true
            this.highlightCells(td)
            const parts = td.id.split('-')
            this.fromDateParts = parts
            this.unitId = parts[0];
            this.fromDate = parts[1] + '-' + parts[2] + '-' + parts[3]
            td.classList.add('selected')
            this.selectedCells.push(td)
            this.bookingId = ''
        } else if (this.toDate === null) {
            this.makingSelection = false
            this.clear()

        } else {
            this.makingSelection = false
            this.selectionMade = true
            td.classList.remove('selected')
        }

    },

    hoverDate: function (td) {

        if (this.bookingSelected) return

        if (this.selectionMade) return

        if (!this.makingSelection && td.classList.contains('booked')) {
            this.showBooking(td)
            this.highlightCells(td)
            return
        }

        const parts = td.id.split('-')
        unitId = parts[0];
        this.highlightCells(td)

        if (!this.makingSelection) {
            this.bookingId = ''
            this.name = ''
            this.guests = 1
            this.unitId = unitId
            this.toDate = null
            this.fromDate = parts[1] + '-' + parts[2] + '-' + parts[3]
            return
        }

        if (this.unitId !== unitId) {
            this.clear()
            return
        }

        if (Number(this.fromDateParts[4]) >= Number(parts[4])) {
            this.clear()
            return
        }

        if (td.classList.contains('selected')) {
            this.lastCell.classList.remove('selected')
            this.toDate = parts[1] + '-' + parts[2] + '-' + parts[3]
            if (this.toDate == this.fromDate) this.toDate = null
            this.lastCell = td
            return
        }

        if (this.lastCell !== null && this.lastCell.classList.contains('booked')) {
            this.clear()
            return
        }

        if (td.classList.contains('booked')) {
            td.classList.add('selected')
            this.toDate = parts[1] + '-' + parts[2] + '-' + parts[3]
            this.lastCell = td
            this.selectedCells.push(td)
            return
        }

        if (!td.classList.contains('selected')) {
            this.lastCell = td
            td.classList.add('selected')
            this.toDate = parts[1] + '-' + parts[2] + '-' + parts[3]
            this.selectedCells.push(td)
            return
        }
    },

    clear: function () {
        this.selectedCells.forEach(cell => {
            cell.classList.remove('selected')
        })
        this.selectionMade = false
        this.makingSelection = false
        this.bookingSelected = false
        this.fromDateParts = null
        this.fromDate = null
        this.toDate = null
        this.unitId = null
        this.selectedCells = []
        this.lastCell = null
        this.bookingId = ''
    },

    highlightCells: function (td) {
        if (this.highlitCells.length > 0 && this.highlitCells[0].dataset.id !== td.dataset.id) {
            this.highlitCells.forEach(cell => {
                cell.style.backgroundColor = this.originalColor
            })
            this.highlitCells = []
        }

        if (this.makingSelection) return

        let id = td.dataset.id;
        let cells = document.querySelectorAll(`td[data-id='${id}']`);
        cells.forEach(cell => {
            if (this.highlitCells.length === 0) {
                this.originalColor = cell.style.backgroundColor
            }
            cell.style.backgroundColor = '#44cc66'
            this.highlitCells.push(cell)
        })
    },

    showBooking(td) {
        this.bookingId = td.dataset.id
        this.name = td.dataset.name
        this.guests = td.dataset.guests
        this.unitId = td.dataset.unit
        this.fromDate = td.dataset.from
        this.toDate = td.dataset.to
    },

    postBooking: function () {
        if (this.fromDate === null || this.toDate === null) return
        const data = {
            name: this.name,
            unitId: Number(this.unitId),
            from: this.fromDate,
            to: this.toDate,
            guests: Number(this.guests)
        }
        fetch('http://localhost:8080/booking', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        }).then(response => {
            if (response.ok) {
                this.clear()
                window.location.reload()
            } else {
                alert('Booking failed: ' + response.status)
            }
        })
    },

    deleteBooking: function () {
        fetch('http://localhost:8080/booking/' + this.bookingId, {
            method: 'DELETE'
        }).then(response => {
            if (response.ok) {
                this.clear()
                window.location.reload()
            } else {
                alert('Booking failed: ' + response.status)
            }
        })
    }

}