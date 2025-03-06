database = [
    {
        alasan : "ppppp",
        date : "12/11/20225"
    },
    {
        alasan : "lupa sahur",
        date : "12/11/20225"
    },
    {
        alasan : "ppppp",
        date : "12/11/20225"
    }
]

const root = document.getElementById("root")
let element_html = "";

database.forEach((element , index) => {
    element_html += `
                <div class="min-h-20 flex  ${(index != database.length - 1) ? "border-b-neutral-700 border-b-2" : ""}">
                    <div class="w-20 justify-center h-20 flex items-center">
                        <input id="select" class="w-full  h-[30%] " type="checkbox" name="" id="">
                    </div>
                    <div class="flex-1 flex items-center">
                        <p class="text-gray-400">${element.alasan}</p>
                    </div>

                    <div class="w-40 flex items-center">
                        <p class="cursor-pointer  text-gray-400">${element.date}</p>
                    </div>
                    <div class="w-30 flex">
                        <button class="cursor-pointer text-red-400">Delete</button>
                    </div>
                </div>
    `
});

root.innerHTML = element_html