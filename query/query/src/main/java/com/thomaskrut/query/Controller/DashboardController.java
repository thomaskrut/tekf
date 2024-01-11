package com.thomaskrut.query.Controller;

import com.thomaskrut.query.Model.DashboardModel;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;

import java.util.concurrent.ExecutionException;

@Controller
public class DashboardController {

    private DashboardModel dashboardModel;

    public DashboardController(DashboardModel dashboardModel) {
        this.dashboardModel = dashboardModel;
    }

    @RequestMapping("/dashboard")
    public String calendar(Model model) throws ExecutionException, InterruptedException {

        dashboardModel.update();

        model.addAttribute("units", dashboardModel.getUnits());
        model.addAttribute("model", dashboardModel);

        return "dashboard.html";
    }

}
