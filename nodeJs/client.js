console.log(">>> Run client")
const net = require('net');
const readline = require('readline');

const HOST = '127.0.0.1';
const PORT = 65456;

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function main() {
    const clientSocket = new net.Socket();

    clientSocket.connect(PORT, HOST, () => {
        console.log('> echo-client is activated');
    });

    clientSocket.on('error', (error) => {
        console.error('> connect() failed by error:', error);
        clientSocket.destroy();
    });

    clientSocket.on('connect', () => {
        rl.question('> ', (sendMsg) => {
            clientSocket.write(sendMsg);
        });
    });

    clientSocket.on('data', (recvData) => {
        console.log('> received:', recvData.toString('utf-8'));
        if (recvData.toString('utf-8') === 'quit') {
            clientSocket.end();
            rl.close();
            console.log('> echo-client is de-activated');
        } else {
            rl.question('> ', (sendMsg) => {
                clientSocket.write(sendMsg);
            });
        }
    });
}

main();
