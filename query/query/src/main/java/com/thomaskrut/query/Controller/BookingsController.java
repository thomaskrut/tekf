package com.thomaskrut.query.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import com.thomaskrut.query.Model.BookingsModel;

@Controller
public class BookingsController {

    private BookingsModel bookingsModel;

    public BookingsController(BookingsModel bookings) {
        this.bookingsModel = bookings;
    }
    @RequestMapping("/bookings")
    public String bookings(Model model) {
        model.addAttribute("bookings", bookingsModel.getBookings());
        return "bookings.html";
    }

}
