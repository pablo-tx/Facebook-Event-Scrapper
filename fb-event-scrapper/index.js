const puppeteer = require('puppeteer');

async function run() {
    const url = process.argv.slice(2);
    const browser = await puppeteer.launch({
        //headless:false
    });
    const page = await browser.newPage();

    const EVENT_NAME = "#upcoming_events_card > div:nth-child(1) > div:nth-child(INDEX) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(2) > div:nth-child(1) > div:nth-child(1) > a:nth-child(1) > span:nth-child(1)";
    const EVENT_MONTH = "#upcoming_events_card > div:nth-child(1) > div:nth-child(INDEX) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > span:nth-child(1) > span:nth-child(1)";
    const EVENT_DAY = "#upcoming_events_card > div:nth-child(1) > div:nth-child(INDEX) > table:nth-child(1) > tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1) > span:nth-child(1) > span:nth-child(2)";
    
    const EVENT_CONTAINER = "#upcoming_events_card > div:nth-child(1)"

    await page.goto(url+"events");
    //await page.waitFor(2000);

    let eventList = await page.evaluate((sel) => {
        return document.querySelectorAll(sel).length != 0 ? document.querySelectorAll(sel)[0].childNodes.length : 0;
    }, EVENT_CONTAINER);

    for (let i = 2; i <= eventList; i++) {
        let eventNameSelector = EVENT_NAME.replace("INDEX", i);
        let eventMonthSelector = EVENT_MONTH.replace("INDEX", i);
        let eventDaySelector = EVENT_DAY.replace("INDEX", i);

        let eventName = await page.evaluate((sel) => {
            let element = document.querySelector(sel);
            return element ? element.innerHTML : null;
        }, eventNameSelector);

        let eventMonth = await page.evaluate((sel) => {
            let element = document.querySelector(sel);
            return element ? element.innerHTML : null;
        }, eventMonthSelector);

        let eventDay = await page.evaluate((sel) => {
            let element = document.querySelector(sel);
            return element ? element.innerHTML : null;
        }, eventDaySelector);

        console.log(eventName+" "+eventMonth+" "+eventDay);
    }

    browser.close();
}

run();