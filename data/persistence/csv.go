package persistence

import (
    "encoding/csv"
    "github.com/jszwec/csvutil"
    log "github.com/sirupsen/logrus"
    "io"
    "os"
    
    "github.com/ul-gaul/go-basestation/data/packet"
)

type CsvPacketReader struct {
    *csvutil.Decoder
}

func NewCsvPacketReader(reader io.Reader) *CsvPacketReader {
    decoder, err := csvutil.NewDecoder(csv.NewReader(reader))
    if err != nil {
        log.Panicln(err)
    }
    return &CsvPacketReader{decoder}
}

func (cr *CsvPacketReader) Read() (packet.RocketPacket, error) {
    var pkt packet.RocketPacket
    err := cr.Decode(&pkt)
    return pkt, err
}

func (cr *CsvPacketReader) ReadAll() ([]packet.RocketPacket, error) {
    var packets []packet.RocketPacket
    err := cr.Decode(&packets)
    return packets, err
}

func ReadCsv(csvPath string) ([]packet.RocketPacket, error) {
    file, err := os.Open(csvPath)
    if err != nil { return nil, err }
    return NewCsvPacketReader(file).ReadAll()
}

/********************************* CSV Writer *********************************/

type CsvPacketWriter struct {
    *csvutil.Encoder
}

func NewCsvPacketWriter(writer io.Writer) *CsvPacketWriter {
    return &CsvPacketWriter{csvutil.NewEncoder(csv.NewWriter(writer))}
}

func (cw *CsvPacketWriter) Write(pkt packet.RocketPacket) error {
    return cw.Encode(pkt)
}

func (cw *CsvPacketWriter) WriteAll(packets []packet.RocketPacket) error {
    return cw.Encode(packets)
}