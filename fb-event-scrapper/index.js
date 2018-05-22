const puppeteer = require('puppeteer');

async function run() {
    const browser = await puppeteer.launch({headless:false});
    const page = await browser.newPage();
    const EVENT_NAME = "#js_5y > span:nth-child(1)";

    await page.goto("https://www.facebook.com/MaeWestGranadaOfficialPage/events");
    await page.waitFor(2 * 1000);

    browser.close();
}

run();