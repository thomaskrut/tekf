package com.thomaskrut.query.Controller;

import com.thomaskrut.query.Model.DashboardModel;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.concurrent.ExecutionException;

@Controller
public class DashboardController {

    private DashboardModel dashboardModel;
    private DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd");

    public DashboardController(DashboardModel dashboardModel) {
        this.dashboardModel = dashboardModel;
    }

    @RequestMapping("/dashboard")
    public String calendar(Model model, @RequestParam(defaultValue = "") String date)
            throws ExecutionException, InterruptedException {

        LocalDate dateToModel;

        switch (date) {
            case "" -> dateToModel = LocalDate.of(LocalDate.now().getYear(), LocalDate.now().getMonth(),
                    LocalDate.now().getDayOfMonth());
            default -> dateToModel = LocalDate.parse(date);
        }

        boolean showButtons = dateToModel.equals(LocalDate.of(LocalDate.now().getYear(), LocalDate.now().getMonth(),
                LocalDate.now().getDayOfMonth()));

        dashboardModel.update(dateToModel);

        model.addAttribute("showButtons", showButtons);
        model.addAttribute("tomorrow", dateToModel.plusDays(1).format(formatter));
        model.addAttribute("yesterday", dateToModel.minusDays(1).format(formatter));
        model.addAttribute("today", dateToModel.format(formatter));
        model.addAttribute("units", dashboardModel.getUnits());
        model.addAttribute("model", dashboardModel);

        return "dashboard.html";
    }

}
