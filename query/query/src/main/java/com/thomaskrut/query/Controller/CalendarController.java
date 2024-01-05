package com.thomaskrut.query.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import com.thomaskrut.query.Model.CalendarModel;

@Controller
public class CalendarController {

    private CalendarModel calendarModel;

    public CalendarController(CalendarModel calendarModel) {
        this.calendarModel = calendarModel;
    }
    @RequestMapping("/calendar")
    public String calendar(Model model) {
        model.addAttribute("calendar", calendarModel.getBookings());
        return "calendar.html";
    }

}
