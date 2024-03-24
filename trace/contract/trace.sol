// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract Teatrace {
    mapping(string => TraceInfo) public tracedata;
    struct TraceInfo {
        TeaInfo tea;
        processInfo process;
        productionInfo production;
        storageInfo store;
        logisInfo logis;
    }

    struct TeaInfo {
        string planet;
        string location;
        string grower;
        string species;
        string quarantine;
        string packages;
        string pick;
        string picktime;
        string picker;
        uint256 inserttime;
        uint256 updatetime;
    }

    struct processInfo {
        string process;
        string processtime;
        string processer;
        string company;
        uint256 inserttime;
        uint256 updatetime;
    }

    struct productionInfo {
        string production;
        string productiontime;
        string producter;
        string company;
        uint256 inserttime;
        uint256 updatetime;
    }

    struct storageInfo {
        string store;
        string storetime;
        string company;
        uint256 inserttime;
        uint256 updatetime;
    }

    struct logisInfo {
        string path;
        string logisway;
        string driver;
        string logistime;
        string company;
        uint256 inserttime;
        uint256 updatetime;
    }

    address producer;
    address process;
    address storer;
    address logostic;
    address sales;
    address tea;

    constructor(
        address _producer,
        address _storer,
        address _logostic,
        address _process,
        address _tea
    ) {
        producer = _producer;
        storer = _storer;
        logostic = _logostic;
        process = _process;
        tea = _tea;
    }

    modifier onlyTea() {
        require(msg.sender == tea, "Only tea can do");
        _;
    }
    modifier onlyProducer() {
        require(msg.sender == producer, "Only producer can do");
        _;
    }
    modifier onlyProcess() {
        require(msg.sender == process, "Only process can do");
        _;
    }
    modifier onlyStore() {
        require(msg.sender == storer, "Only storer can do");
        _;
    }
    modifier onlyLogis() {
        require(msg.sender == logostic, "Only logostic can do");
        _;
    }

    function addOrupdateTeaData(
        string memory id,
        string memory planet,
        string memory location,
        string memory grower,
        string memory species,
        string memory quarantine,
        string memory packages,
        string memory pick,
        string memory picktime,
        string memory picker
    ) external onlyTea {
        uint256 inserttime;
        if (tracedata[id].tea.inserttime == 0) {
            inserttime = block.timestamp;
        }
        tracedata[id].tea = TeaInfo({
            planet: planet,
            location: location,
            grower: grower,
            species: species,
            quarantine: quarantine,
            packages: packages,
            pick: pick,
            picktime: picktime,
            picker: picker,
            inserttime: inserttime,
            updatetime: block.timestamp
        });
    }

    function addOrupdateProdData(
        string memory id,
        string memory production,
        string memory productiontime,
        string memory producter,
        string memory company
    ) external onlyProducer {
        uint256 inserttime;
        if (tracedata[id].production.inserttime == 0) {
            inserttime = block.timestamp;
        }
        tracedata[id].production = productionInfo({
            production: production,
            productiontime: productiontime,
            producter: producter,
            company: company,
            inserttime: inserttime,
            updatetime: block.timestamp
        });
    }

    function addOrupdateStoreData(
        string memory id,
        string memory store,
        string memory storetime,
        string memory company
    ) external onlyStore {
        uint256 inserttime;
        if (tracedata[id].store.inserttime == 0) {
            inserttime = block.timestamp;
        }
        tracedata[id].store = storageInfo({
            store: store,
            storetime: storetime,
            company: company,
            inserttime: inserttime,
            updatetime: block.timestamp
        });
    }

    function addOrupdateProcessData(
        string memory id,
        string memory processa,
        string memory processtime,
        string memory processer,
        string memory company
    ) external onlyProcess {
        uint256 inserttime;
        if (tracedata[id].process.inserttime == 0) {
            inserttime = block.timestamp;
        }
        tracedata[id].process = processInfo({
            process: processa,
            processtime: processtime,
            processer: processer,
            company: company,
            inserttime: inserttime,
            updatetime: block.timestamp
        });
    }

    function addOrupdateLogissData(
        string memory id,
        string memory path,
        string memory logisway,
        string memory logistime,
        string memory driver,
        string memory company
    ) external onlyLogis {
        uint256 inserttime;
        if (tracedata[id].logis.inserttime == 0) {
            inserttime = block.timestamp;
        }
        tracedata[id].logis = logisInfo({
            path: path,
            logisway: logisway,
            driver: driver,
            logistime: logistime,
            company: company,
            inserttime: inserttime,
            updatetime: block.timestamp
        });
    }
}
