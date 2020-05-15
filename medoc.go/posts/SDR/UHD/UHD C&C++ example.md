```
{
    "url": "UHD_C&C++_example",
    "time": "2018/01/22 19:45",
    "tag": "SRD, UHD, C&C++"
}
```


# UHD C / C ++编程实例USRP发送，接收数据

## 1. UHD库函数简介

### 1.1发送函数

##### 新建一个USRP设备

```
std::string args = " "; args指定USRP地址，配置USRP参数等
uhd::usrp::multi_usrp::sptr usrp = uhd::usrp::multi_usrp::make(args);
```

##### 设置时钟源和时钟频率

```
std::string ref = "internal"; // internal, external, MIMO.
usrp->set_clock_source(ref);
double rate = 40e6;
usrp->set_master_clock_rate(rate);

usrp->get_clock_source() //获取当前的时钟源
usrp->get_clock_sources() //获取所有可用的时钟源
usrp->get_master_clock_rate(); //获取时钟频率
```

##### 设置采样率

```
double samp_rate = 20e6;
usrp->set_tx_rate(samp_rate);

usrp->get_tx_rate() //获取当前的采样率
usrp->get_tx_rates() //获取采样率的采样范围
```

##### 设置发射中心频点

```
double freq = 2.412e9;
usrp->set_tx_freq(freq);

usrp->get_fe_tx_freq_range();
usrp->get_tx_freq();
usrp->get_tx__freq_range();
```

##### 设置发射增益

```
double tx_gain = 60;
usrp->set_tx_gain(tx_gain);
```

##### 发射设置天线

```
默认情况下无需设置发射天线，如需指定特定发射天线时设置。
usrp->set_tx_antenna(0);
usrp->get_tx_antenna();
usrp->get_tx_antennas();
```

##### 创建发送流

##### 新建发送流参数

```
std::string cpu_format = "fc32";
目前cpu_format支持的类型有：fc64 - 复杂fc32 - 复杂sc16 - 复杂sc8 - 复杂
std::string wire_format = "sc16";
目前wire_format支持的类型有：sc16 - Q16 I16 sc8 - Q8_1 I8_1 Q8_0 I8_0

uhd::stream_args_t stream_args(cpu_format, wire_format);
uhd::tx_streamer::sptr tx_stream = usrp->get_tx_stream(stream_args);
```

##### 中频数据管理

```
uhd::tx_metada_t md; //结构体描述接收到的中频信号
md.start_of_burst = false; // 设置为真的时候发送第一个数据包
md.end_of_burst = false; // 设置为真的时候发送最后一个数据包
md.has_time_spec = false; // 设置为false时立即发送，设置为真的时候在特定时间发送
```

### 1.2接收函数

##### 创建一个USRP

```
std::string args = " ";
uhd::usrp::multi_usrp::sptr usrp = uhd::usrp::multi_usrp::make(args);
```

##### 设置时钟源和时钟频率

```
std::string ref = "internal"; //internal, external and MIMO
usrp->set_clock_source(ref);

double clock_rate = 40e6;
usrp->set_master_clock_rate(clock_rate);
```

##### 设置采样率

```
double rate = 20e6;
usrp->set_rx_rate(rate);
```

##### 设置中心频率

```
double freq = 2.412e9;
usrp->set_rx_freq(freq);
```

##### 设置增益

```
double rx_gain = 50;
usrp->set_rx_gain(rx_gain);
```

##### 创建接收流

```
std::string cpu_format = "fc32";
std::string wire_format = "sc16";
uhd::stream_args_t stream_args(cpu_format, wire_format);
uhd::rx_streamer::sptr rx_stream = usrp->get_rx_stream(stream_args);
```

##### 设置接收模式

```
uhd::stream_cmd_t stream_cmd(uhd::stream_cmd_t::STREAM_MODE_START_CONTINUOUS);

//UHD_STREAM_MODE_START_CONTINUOUS=97, //UHD_STREAM_MODE_STOP_CONTINUOUS=111, //UHD_STREAM_MODE_NUM_SAMPS_AND_DONE=100, //UHD_STREAM_MODE_NUM_SAMPS_AND_MORE=109

stream_cmd.num_samps = 1000; // 接收采样点数
stream_cmd.stream_now = true; //现在开始接收
rx_stream->issue_stream_cmd(stream_cmd); //配置rx_stream参数
```

### 2. USRP发送数据
　　代码实现USRP发送二进制文件802.11a_BUPT41.seg内部数据。里面的数据为float复数类型，复数的实部和虚部交替存储。即如下图所示。用户可以把要发送的数据生成对应的二进制文件，也可以直接写成数组放在代码中。注意：float类型数据的单位1就是“1”。
这里写图片描述


源码

```
#include <uhd/usrp/multi_usrp.hpp>
#include <signal.h>

#define SAMPLE_PER_BUFF 2000

int stop_signal_called = 0;
void sig_int_handle(int)
{
    stop_signal_called = 1;
    printf("stop tx.........\n");
    exit(0);
}

int main()
{
    std::string addr_args = " ";
    uhd::usrp::multi_usrp::sptr usrp = uhd::usrp::multi_usrp::make(addr_args);
    printf("Create a usrp......\n");

    // set the ref and clock rate
    std::string ref = "internal";
    usrp->set_clock_source(ref);

    float clock_rate = 40e6;
    usrp->set_master_clock_rate(clock_rate);
    printf("set the  clock rate %0.2f \n", usrp->get_master_clock_rate() );

    // set the sample rate
    float samp_rate = 20e6;
    usrp->set_tx_rate(samp_rate);
    printf("set the tx sample rate to %0.2f \n", usrp->get_tx_rate());

    // set the center frequency
    float center_freq = 2.412e9;
    usrp->set_tx_freq(center_freq);
    printf("set the tx center freq to %0.2f \n", usrp->get_tx_freq());

    // set the rf gain
    float tx_gain = 90;
    usrp->set_tx_gain(tx_gain);
    printf("set the tx gain to %0.2f \n", usrp->get_tx_gain());


    // create a tx stream
    std::string cpu_format = "fc32";
    std::string wire_format = "sc16";
    uhd::stream_args_t stream_args(cpu_format, wire_format);
    uhd::tx_streamer::sptr tx_stream = usrp->get_tx_stream(stream_args);

    uhd::tx_metadata_t md;

    // catch the INT signal
    signal(SIGINT, sig_int_handle);
    float read_buff[SAMPLE_PER_BUFF * 2] = {0};

    while(!stop_signal_called)
    {
        FILE *fp = fopen("802.11a_BUPT41.seg", "rb");
        md.start_of_burst = false;
        md.end_of_burst = false;


        while( (!md.end_of_burst) && (!stop_signal_called) )
        {
            int read_length = 0;
            if( (read_length = fread(read_buff, sizeof(uint32_t), SAMPLE_PER_BUFF * 2, fp) ) == (SAMPLE_PER_BUFF * 2) )
            {
                //int index;
                //for(index = 0; index < SAMPLE_PER_BUFF * 2; index++)
                //    printf("%0.2f ", read_buff[index]);
                //puts("");

                //md.start_of_burst = true;
                tx_stream->send(read_buff, SAMPLE_PER_BUFF, md);
                //md.start_of_burst = false;
                //sleep(1);
            }
            else if(read_length >= 0)
            {
                md.end_of_burst = true;
            }
        }

        fclose(fp);
    }


    return 0;
}
```


g ++编译时需要链接uhd库，系统需要安装UHD驱动。具体过程参见之前博客。Ubuntu14.04源码安装UHD3.8.0

```
g++ tx.cpp -o tx -luhd
```

### 3. USRP接收数据
　　启动USRP，并使USRP一直接收数据。



```
#include <uhd/usrp/multi_usrp.hpp>
#include <csignal>

#define SAMPLE_PER_BUFF 2000

int stop_signal_called = false;
void sig_int_handle()
{
    stop_signal_called = true;
}

int main()
{
    std::string addr_args = " ";
    uhd::usrp::multi_usrp::sptr usrp = uhd::usrp::multi_usrp::make(addr_args);

    // set the clock source and clock rate
    std::string ref = "internal";
    usrp->set_clock_source(ref);

    float clock_rate = 40e6;
    usrp->set_master_clock_rate(clock_rate);
    printf("set the  clock rate %0.2f \n", usrp->get_master_clock_rate() );

    // set the sample rate
    float samp_rate = 20e6;
    usrp->set_rx_rate(samp_rate);
    printf("set the tx sample rate to %0.2f \n", usrp->get_rx_rate());

    // set the center frequency
    float center_freq = 2.412e9;
    usrp->set_rx_freq(center_freq);
    printf("set the tx center freq to %0.2f \n", usrp->get_rx_freq());

    // set the rf gain
    float rx_gain = 80;
    usrp->set_rx_gain(rx_gain);
    printf("set the tx gain to %0.2f \n", usrp->get_rx_gain());

    std::string cpu_format = "fc32";
    std::string wire_format = "sc16";
    uhd::stream_args_t stream_args(cpu_format, wire_format);
    uhd::rx_streamer::sptr rx_stream = usrp->get_rx_stream(stream_args);

    uhd::rx_metadata_t md;

    //uhd::stream_cmd_t stream_cmd(uhd::stream_cmd_t::STREAM_MODE_NUM_SAMPS_AND_DONE);
    uhd::stream_cmd_t stream_cmd(uhd::stream_cmd_t::STREAM_MODE_START_CONTINUOUS);
    stream_cmd.num_samps = SAMPLE_PER_BUFF;
    stream_cmd.stream_now = true;
    //stream_cmd.time_spec = uhd::time_spec_t();
    stream_cmd.time_spec = usrp->get_time_now();
    rx_stream->issue_stream_cmd(stream_cmd);

    uint32_t buff[SAMPLE_PER_BUFF*2] = {0};
    unsigned long long num_total_samps = 0;

    while(!stop_signal_called)
    {
        //int num_rx_samps = rx_stream->recv(buff, SAMPLE_PER_BUFF, md);
        memset(buff, 0, SAMPLE_PER_BUFF * sizeof(uint32_t));
        int num_rx_samps = rx_stream->recv(buff, SAMPLE_PER_BUFF, md, 3.0, false);
        if(md.error_code == uhd::rx_metadata_t::ERROR_CODE_TIMEOUT)
        {
            printf("Timeout while streaming......\n");
            break;
        }
        if(md.error_code == uhd::rx_metadata_t::ERROR_CODE_OVERFLOW)
        {
            printf("Overflowing while stream......\n");
            continue;
        }
        if(md.error_code != uhd::rx_metadata_t::ERROR_CODE_NONE)
        {
            //printf("Receive error: %s \n", md.strerror());
            continue;
        }

    //printf("num_rx_samps = %d \n",num_rx_samps);
        num_total_samps += num_rx_samps;

    }
    printf("num_total_samps = %d \n", num_total_samps);

    stream_cmd.stream_mode = uhd::stream_cmd_t::STREAM_MODE_STOP_CONTINUOUS;
    rx_stream->issue_stream_cmd(stream_cmd);

    return 0;
}
```


g ++编译时需要链接uhd库，系统需要安装UHD驱动。具体过程参见之前博客。Ubuntu14.04源码安装UHD3.8.0

```
g++ rx.cpp -o rx -luhd
```

### 4。结论
　　以上代码在Ubuntu上采用USRP B200测试通过。
　　掌握UHD库函数使用，g ++链接UHD库等是基于USRP开发软件无线电项目的基础。之后，我们可以进一步设计发送程序和接收程序，用USRP搭建一套SDR收发系统。
