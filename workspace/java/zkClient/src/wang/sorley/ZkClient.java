package wang.sorley;

import org.apache.zookeeper.ZooKeeper;
import org.apache.zookeeper.admin.ZooKeeperAdmin;
import java.util.ArrayList;
import java.util.List;

public class ZkClient {
    public static void main(String[] args) {
        ZooKeeper zk = new ZooKeeperAdmin();

        List<String> newMembers = new ArrayList<String>();
        newMembers.add("server.1=1111:1234:1235;1236");
        newMembers.add("server.2=1112:1237:1238;1239");
        newMembers.add("server.3=1114:1240:1241:observer;1242");

        byte[] config = zk.reconfig(null, null, newMembers, -1, new Stat());

        String configStr = new String(config);
        System.out.println(configStr);
    }
}
