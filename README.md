# its_manager

전달 받은 이미지 변환 요청 성격에 따라 이미지를 변환 하고 Upload를 수행한다.

## proc_stream
> Kafka 로 부터 imglog 를 Consume 하고 분석한다.

## proc_convert
> ImagemagicK를 사용하여 필요시 resize를 수행한다.

## proc_archive
> 요청 받은 압축방식에 따라 이미지 압축을 수행한다.


