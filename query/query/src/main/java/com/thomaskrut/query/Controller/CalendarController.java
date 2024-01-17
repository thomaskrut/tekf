package com.thomaskrut.query.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import com.thomaskrut.query.Model.CalendarModel;

import java.util.concurrent.ExecutionException;

@Controller
public class CalendarController {

    private CalendarModel calendarModel;

    public CalendarController(CalendarModel calendarModel) {
        this.calendarModel = calendarModel;
    }
    @RequestMapping("/calendar")
    public String calendar(Model model) throws ExecutionException, InterruptedException {

        calendarModel.update();
        model.addAttribute("days", calendarModel.getCalendar().getDays());
        model.addAttribute("units", calendarModel.getCalendar().getUnits());
        return "calendar.html";
    }

}
