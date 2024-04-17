console.log(">>> Run server")
const net = require('net');

const HOST = '127.0.0.1';
const PORT = 65456;

function main() {
    const serverSocket = net.createServer();

    serverSocket.on('listening', () => {
        console.log('> echo-server is activated');
    });

    serverSocket.on('error', (error) => {
        console.error('> Server error:', error);
        serverSocket.close();
    });

    serverSocket.on('connection', (clientSocket) => {
        console.log('> client connected by IP address', clientSocket.remoteAddress, 'with Port number', clientSocket.remotePort);

        clientSocket.on('data', (recvData) => {
            console.log('> echoed:', recvData.toString('utf-8'));
            clientSocket.write(recvData);
            if (recvData.toString('utf-8') === 'quit') {
                clientSocket.end();
            }
        });

        clientSocket.on('error', (error) => {
            console.error('> Client socket error:', error);
        });

        clientSocket.on('end', () => {
            console.log('> Client disconnected');
        });
    });

    serverSocket.listen(PORT, HOST);
}

main();
