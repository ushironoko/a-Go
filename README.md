# a-Go

超A&Gの番組を rtmpdump を使って録画できるシンプルなツール。顎。

# 使い方

実行ファイルはリリースからダウンロードするか、本コードをビルドしてください。（ビルドにはGo の実行環境が必要です）

録画に rtmpdump を使用します。Windows の場合は以下のサイトから windows.zip をダウンロードして解凍後、rtmpdump.exe を main.exe と同じ場所に配置してください。

http://rtmpdump.mplayerhq.hu/download/

Linux または Mac の場合は コマンド等で rtmpdump をインストールしてください。

タスクスケジューラや crontab で定期的に実行すれば、自動録音環境を作る事もできます。

# 設定

setting.json に出力ファイルに関する設定を記述できます。

* startHour(現在時刻と比較されて、一致した場合この設定を利用します) 
* recTime(録画時間を分単位で設定します)
* outputDir(録画ファイルをどこに出力するか指定します,デフォルトはdist)
* rtmpUrl(接続するRTMP URLを指定します)
* fileName(出力されるファイルの名前です)
* fileExtention(出力されるファイルの拡張子です,現在flvのみ対応)

# 動作確認

* Windows 10 (タスクスケジューラ)
* Linux (Ubuntu 18)

# 注意

録画したファイルの再配布に関しては法律で禁止されています。私用の範囲でご利用ください。

上記に関して一切の責任を負いかねますのでご了承ください。