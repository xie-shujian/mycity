function cityPage(CurrentPage) {
    fetch("/city/data?page=" + CurrentPage)
        .then(response => response.json())
        .then(data => {
            //   console.log(data);
            let mytrs = "";
            for (let i = 0; i < data.cities.length; i++) {
                let idx = i + 1 + (data.page.current_page - 1) * data.page.page_size;
                mytrs += "<tr>"
                    + "<td>" + idx + "</td>"
                    + "<td>" + data.cities[i].id + "</td>"
                    + "<td>" + data.cities[i].name + "</td>"
                    + "<td>" + data.cities[i].countryCode + "</td>"
                    + "<td>" + data.cities[i].district + "</td>"
                    + "<td>" + data.cities[i].population + "</td>"
                    + "</tr>"
                    ;
            }
            document.querySelector("#dataList").innerHTML = mytrs;
            createPageNav(data.page)
        });
}

function createPageNav(page) {
    //let navSize = 8;
    let step = 3; //2*2*2=8
    let navPage = (page.current_page - 1 >> step) + 1;
    let navMin = ((navPage - 1) << step) + 1;
    let navMax = navPage << step;
    navMax = Math.min(navMax, page.total_pages);

    let liPrevious = "<li class=\"page-item\"><a class=\"page-link\" href=\"#\">Previous</a></li>";
    if (page.current_page == 1) {
        liPrevious = "<li class=\"page-item disabled\"><a class=\"page-link \">Previous</a></li>";
    }
    let mylis = liPrevious;

    for (let i = navMin; i <= navMax; i++) {
        let myli = "<li class=\"page-item\">"
            + "<a class=\"page-link\" href=\"#\" >" + i + "</a>"
            + "</li>";
        if (i == page.current_page) {
            myli = "<li class=\"page-item active\" aria-current=\"page\">"
                + "<a class=\"page-link\" href=\"#\" >" + i + "</a>"
                + "</li>";
        }
        mylis += myli;
    }

    let liNext = "<li class=\"page-item\"><a class=\"page-link\" href=\"#\">Next</a></li>";
    if (page.current_page == page.total_pages) {
        liNext = "<li class=\"page-item disabled\"><a class=\"page-link\">Next</a></li>";
    }
    mylis += liNext;

    document.querySelector("#pageList").innerHTML = mylis;
}

function loadFirstPage() {
    cityPage(1);
}
function loadOtherPage(event) {
    //disabled
    if(event.target.classList.contains("disabled")){
        return;
    }
    

    let slastPage = document.querySelector("#pageList").querySelector(".active").textContent;
    let iLastPage = parseInt(slastPage)

    let currentPage = event.target.textContent;
    if (currentPage == "Next") {
        currentPage = String(iLastPage + 1);
    }
    else if (currentPage == "Previous") {
        currentPage = String(iLastPage - 1);
    }
    //console.log(event.target.textContent);

    // console.log(lastPage);

    cityPage(currentPage);
}